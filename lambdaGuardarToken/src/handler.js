'use strict';
const serverless = require('serverless-http');
const express = require('express');
const app = express();
const AWS = require('aws-sdk');
const bodyParser = require('body-parser');

const USERS_TABLE = process.env.USERS_TABLE;
const IS_OFFLINE = process.env.IS_OFFLINE;
let dynamoDB;

if (IS_OFFLINE === 'true') {
  dynamoDB = new AWS.DynamoDB.DocumentClient({
    region: 'localhost',
    endpoint: 'http://localhost:8000'
  });
} else {
  dynamoDB = new AWS.DynamoDB.DocumentClient();
}

app.use((req, res, next) => {
  res.header('Access-Control-Allow-Origin', '*');
  res.header('Access-Control-Allow-Headers', 'Authorization, X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Allow-Request-Method');
  res.header('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, DELETE');
  res.header('Allow', 'GET, POST, OPTIONS, PUT, DELETE');
  next();
});

app.use(bodyParser.urlencoded({ extended: true }));

// app.get('/', (req, res) => {
//   res.send('Hola Mundo con Expressjs');
// });

app.post('/setToken', (req, res) => {
  const { userId, identification, dataToken } = req.body;

  const createdAt = new Date();

  const newToken = {
    userId,
    identification,
    dataToken,
    createdAt,
  }

  const params = {
    TableName: USERS_TABLE,
    Item: newToken
  };

  dynamoDB.put(params, (error) => {
    if (error) {
      res.status(400).json({
        error: 'No se ha podido crear el usuario'
      })
    } else {
      res.json({ userId, identification, dataToken, createdAt });
    }
  });
});

app.put('/setToken', (req, res) => {
  // const { userId, identification, dataToken } = req.body;

  const { userId, dataToken } = req.body;

  const params = {
    TableName: USERS_TABLE,
    Key: { userId },
    UpdateExpression: 'set dataToken = :dataToken',
    ExpressionAttributeValues: {
      ":dataToken": dataToken
    },
    ReturnValues: 'ALL_NEW'
  };

  dynamoDB.update(params, (error, result) => {
    if (error) {
      console.log(error);
      return res.status(400).json({
        error: 'No se ha podido acceder al usuario'
      })
    }
    if (result.Item) {
      const { userId, identification, dataToken, createdAt } = result.Attributes;
      return res.json({ userId, identification, dataToken, createdAt });
    } else {
      res.json({ message: "No encontrado" });
    }
  })
});

app.get('/getToken', (req, res) => {

  const { userId } = req.body;

  const params = {
    TableName: USERS_TABLE,
    Key: {
      userId,
    }
  };

  dynamoDB.get(params, (error, result) => {
    if (error) {
      console.log(error);
      return res.status(400).json({
        error: 'No se ha podido acceder al usuario'
      })
    }
    if (result.Item) {
      const { userId, identification, dataToken, createdAt } = result.Item;
      return res.json({ userId, identification, dataToken, createdAt });
    } else {
      res.status(404).json({ error: 'Usuario no encontrado' })
    }
  })
});

app.get('/getTokenById', (req, res) => {

  const { identification } = req.body;

  const params = {
    TableName: USERS_TABLE,
    IndexName: 'id-identification',
    KeyConditionExpression: "identification = :identification",
    ExpressionAttributeNames: {
      "identification": "identification"
    },
    ExpressionAttributeValues: {
      ":identification": identification
    },
    ReturnValues: 'ALL_NEW'
  };

  dynamoDB.get(params, (error, result) => {
    if (error) {
      console.log(error);
      return res.status(400).json({
        error: 'No se ha podido acceder al usuario'
      })
    }
    if (result.Item) {
      const { userId, identification, dataToken, createdAt } = result.Item;
      return res.json({ userId, identification, dataToken, createdAt });
    } else {
      res.status(404).json({ error: 'Usuario no encontrado' })
    }
  })
});

app.put('/updateToken', (req, res) => {

  const { identification } = req.body;

  const params = {
    TableName: USERS_TABLE,
    Key: { identification },
    UpdateExpression: 'set dataToken = :dataToken',
    ExpressionAttributeValues: {
      ":dataToken": dataToken
    },
    ReturnValues: 'ALL_NEW'
  };

  dynamoDB.update(params, (error, result) => {
    if (error) {
      console.log(error);
      return res.status(400).json({
        error: 'No se ha podido acceder al usuario'
      })
    }
    if (result.Item) {
      const { userId, identification, dataToken, createdAt } = result.Attributes;
      return res.json({ userId, identification, dataToken, createdAt });
    } else {
      res.status(404).json({ error: 'Usuario no encontrado' })
    }
  })
});

module.exports.generic = serverless(app);
