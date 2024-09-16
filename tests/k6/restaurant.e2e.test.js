import http from 'k6/http';
import {check, group} from 'k6';
import errorHandler from './errorHandler.js';

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

const BASE_URL = __ENV || 'localhost:50051';

// Sleep duration between successive requests.
// You might want to edit the value of this variable or remove calls to the sleep function on the script.
const SLEEP_DURATION = 0.1;
// Global variables should be initialized.

export default function() {
  // group('/v1/orders', () => {
  //   {
  //     let url = BASE_URL + `/v1/orders`;
  //     let body = {
  //       'items': [
  //         {
  //           'menu_item_id': '6685d61813c4956eac2592d0',
  //           'quantity': 5,
  //         },
  //         {
  //           'menu_item_id': '6685d61d13c4956eac2592d1',
  //           'quantity': 3,
  //         },
  //       ],
  //       'restaurant_id': '6685d60013c4956eac2592cf',
  //       'user_id': '6685d5d9a1fddcdd0872b0ed',
  //     };
  //     let params = {
  //       headers: {
  //         'Content-Type': 'application/json',
  //         'Accept': 'application/json',
  //       },
  //     };
  //     let request = http.post(url, JSON.stringify(body), params);
  //
  //     errorHandler.logError(!check(request, {
  //       'create an order is ok': (r) => r.status === 200,
  //     }), request);
  //   }
  // });
}
