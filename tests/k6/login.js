import http from 'k6/http';

const url = 'https://blackhorseya.us.auth0.com/oauth/token';

// 从环境变量获取敏感信息
export const adminCredentials = {
  client_id: __ENV.AUTH0_CLIENT_ID || 'your_auth0_client_id',
  client_secret: __ENV.AUTH0_CLIENT_SECRET || 'your_auth0_client_secret',
  username: __ENV.ADMIN_USERNAME || 'admin@gmail.com',
  password: __ENV.ADMIN_PASSWORD || 'um.3YN7m7hRU',
  audience: 'https://godine.seancheng.space/',
  grant_type: 'http://auth0.com/oauth/grant-type/password-realm', // 更新 grant_type
  realm: __ENV.AUTH0_REALM || 'Username-Password-Authentication', // 添加 realm 参数
  scope: 'openid profile email',
};

export function login(credentials) {
  const response = http.post(url, credentials, {
    headers: {'Content-Type': 'application/x-www-form-urlencoded'},
  });

  // 错误处理
  if (response.status !== 200) {
    console.error(`登录失败，状态码：${response.status}，响应：${response.body}`);
    throw new Error('登录失败，无法获取访问令牌');
  }

  const tokenData = response.json();

  if (!tokenData.access_token) {
    console.error('响应中未包含访问令牌');
    throw new Error('登录失败，未获取到访问令牌');
  }

  return tokenData;
}


