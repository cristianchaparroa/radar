#!/bin/bash

npm install npm-force-resolutions --save-dev
npm install
npx npm-force-resolutions
npm install
ionic build
