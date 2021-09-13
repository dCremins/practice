// const express = require('express');
// const cors = require('cors');
import express from 'express';
import cors from 'cors';
import { config } from './config';

const app = express();

app.use(cors());

app.get('/', (req, res) => {
    res.status(200).send('hello world');
});

app.listen(config.port, config.host, (e)=> {
    if(e) {
        throw new Error('Internal Server Error');
    }
    console.log(`${config.name} running on ${config.host}:${config.port}`);
});