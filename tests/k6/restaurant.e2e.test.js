import grpc from 'k6/net/grpc';
import {check, sleep} from 'k6';

const scenarios = {
  average_load: {
    executor: 'ramping-vus',
    stages: [
      {duration: '20s', target: 10},
      {duration: '40s', target: 10},
      {duration: '20s', target: 50},
      {duration: '40s', target: 50},
      {duration: '20s', target: 100},
      {duration: '40s', target: 100},
    ],
  },
  peak_load: {
    executor: 'constant-vus',
    vus: 100,
    duration: '1m',
  },
};

const {SCENARIO} = __ENV;

export const options = {
  cloud: {
    // Project: Default project
    projectID: 3690299,
    // Test runs with the same name groups test runs together.
    name: 'godine restaurant test',
  },

  // define thresholds
  thresholds: {
    http_req_failed: ['rate<0.01'], // http errors should be less than 1%
    http_req_duration: ['p(99)<1000'], // 99% of requests should be below 1s
  },

  // define scenarios
  scenarios: SCENARIO ? {
    [SCENARIO]: scenarios[SCENARIO],
  } : scenarios,
};

const BASE_URL = __ENV.BASE_URL || 'localhost:50051';
const client = new grpc.Client();

// Sleep duration between successive requests.
// You might want to edit the value of this variable or remove calls to the sleep function on the script.
const SLEEP_DURATION = 0.1;
// Global variables should be initialized.

export default function() {
  client.connect(BASE_URL, {reflect: true, plaintext: true});

  // 测试 CreateRestaurant 方法
  let createResponse = client.invoke(
      'restaurant.RestaurantService/CreateRestaurant', {
        name: '测试餐厅',
        address: {
          street: '测试街道',
          city: '测试城市',
          state: '测试省份',
          zip_code: '123456',
        },
      });

  check(createResponse, {
    'CreateRestaurant 成功': (r) => r && r.status === grpc.StatusOK,
    'CreateRestaurant 返回有效的餐厅 ID': (r) => r && r.message &&
        r.message.id !== undefined && r.message.id !== null && r.message.id !==
        '',
  });

  let restaurantId = createResponse.message.id;

  // 测试 GetRestaurant 方法
  let getResponse = client.invoke('restaurant.RestaurantService/GetRestaurant',
      {
        restaurant_id: restaurantId,
      });

  check(getResponse, {
    'GetRestaurant 成功': (r) => r && r.status === grpc.StatusOK,
    'GetRestaurant 返回正确的餐厅 ID': (r) => r && r.message && r.message.id ===
        restaurantId,
  });

  // 测试 ListRestaurantsNonStream 方法
  let listResponse = client.invoke(
      'restaurant.RestaurantService/ListRestaurantsNonStream', {
        page: 1,
        page_size: 10,
      });

  check(listResponse, {
    'ListRestaurantsNonStream 成功': (r) => r && r.status === grpc.StatusOK,
    'ListRestaurantsNonStream 返回餐厅列表': (r) => r && r.message &&
        r.message.restaurants && r.message.restaurants.length > 0,
  });

  // 测试 ListRestaurants 方法（流式响应）
  client.invoke('restaurant.RestaurantService/ListRestaurants', {
    page: 1,
    page_size: 10,
  }, (stream) => {
    stream.on('data', (message) => {
      console.log('接收到餐厅信息:', message);
    });

    stream.on('error', (error) => {
      console.error('流错误:', error);
    });

    stream.on('end', () => {
      console.log('流结束');
    });
  });

  // 等待流式响应完成
  sleep(SLEEP_DURATION);

  client.close();
}
