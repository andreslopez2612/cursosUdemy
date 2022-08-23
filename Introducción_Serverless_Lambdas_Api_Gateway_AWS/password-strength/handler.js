'use strict';
const verifyPasswordLength = require('./constraints/verifyPasswordLength');
module.exports.password = async (event) => {
  try {
    const { password } = event.pathParameters;
    await verifyPasswordLength(password);
    return {
      statusCode: 200,
      body: JSON.stringify({
        message: 'ok',
        password
      })
    }
  } catch (e) {
    return {
      statusCode: 400,
      body: JSON.stringify({
        message: e.message
      })
    }
  }
};
