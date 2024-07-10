/*
 * Godine User Restful API
 * Godine User Restful API document.
 *
 * OpenAPI spec version: 0.1.0
 * Contact: blackhorseya@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://github.com/OpenAPITools/openapi-generator
 *
 * Generator version: 7.7.0-SNAPSHOT
 */

import http from 'k6/http';
import {check, group} from 'k6';
import errorHandler from './errorHandler.js';

const BASE_URL = 'http://localhost:1994/api';

export default function() {
  group('User API CRUD Operations', () => {
    let createdUserId = null;

    // Create User
    let url = BASE_URL + `/v1/users`;
    let body = {
      'address': {
        'city': 'string',
        'state': 'string',
        'street': 'string',
        'zipCode': 'string',
      },
      'email': 'string',
      'name': 'string',
      'password': 'string',
    };
    let params = {
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
    };
    let request = http.post(url, JSON.stringify(body), params);

    if (check(request, {'create an user': (r) => r.status === 201})) {
      createdUserId = request.json('data.id');
      console.log('User created successfully with ID: ' + createdUserId);
    } else {
      console.log(`Unable to create a user ${request.status} ${request.body}`);
      return; // Exit if user creation failed
    }

    // Get User by ID
    let id = createdUserId;
    if (!id) {
      console.log('No user ID available to fetch.');
      return;
    }

    url = BASE_URL + `/v1/users/${id}`;
    request = http.get(url);

    errorHandler.logError(!check(request, {
      'get user by id': (r) => r.status === 200,
    }), request);

    // Get Users List
    let size = '10';
    let page = '1';

    url = BASE_URL + `/v1/users?page=${page}&size=${size}`;
    request = http.get(url);

    errorHandler.logError(!check(request, {
      'get user list': (r) => r.status === 200,
    }), request);

    // Update User Status
    url = BASE_URL + `/v1/users/${id}/status`;
    body = {'is_active': true};
    params = {
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
    };
    request = http.patch(url, JSON.stringify(body), params);

    errorHandler.logError(!check(request, {
      'update user status': (r) => r.status === 200,
    }), request);

    // Delete User by ID
    url = BASE_URL + `/v1/users/${id}`;
    request = http.del(url);

    errorHandler.logError(!check(request, {
      'delete user by id': (r) => r.status === 204,
    }), request);
  });
}