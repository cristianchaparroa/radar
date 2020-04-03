# Radar

## 1. Local Deployment

### Install dependences

```
chmod +x ./bin/install.sh
./bin/install.sh
```


## 2. Documentation
### Background geolocation:

The following are resources according with the Ionic Background geolocation

https://ionicframework.com/docs/native/background-geolocation
https://github.com/ionic-team/ionic-native/issues/2893

## 3. Known issues

### Proposal-numeric-separator

It is a common issue that appears when is run the command `ionic build`

```
An unhandled exception occurred: [BABEL] /Users/cristianchaparroa/radar/web/radar/www/focus-visible-15ada7f7-js-es2015.js: Could not find plugin "proposal-numeric-separator". Ensure there is an entry in ./available-plugins.js for it. (While processing: "/Users/cristianchaparroa/radar/web/radar/node_modules/@babel/preset-env/lib/index.js")
See "/private/var/folders/kj/9h016pgx243cq8f5k2gs32n40000gn/T/ng-YJXCPE/angular-errors.log" for further details.
```
The fix is the following:

```
sudo rm -r node_modules package-lock.json
# add "resolutions": { "@babel/preset-env": "^7.8.7" } to package.json
npm install npm-force-resolutions --save-dev
npm install
npx npm-force-resolutions
npm install
ionic build
```


Reference: https://forum.ionicframework.com/t/could-not-find-plugin-proposal-numeric-separator/185556/6
