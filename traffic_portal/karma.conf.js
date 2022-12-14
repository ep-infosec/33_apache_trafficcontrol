/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Karma configuration file, see link for more information
// https://karma-runner.github.io/1.0/config/configuration-file.html

module.exports = function (config) {
    config.set({
        basePath: "",
        files: [
            "app/dist/public/resources/assets/js/shared-libs.js",
            "app/dist/public/resources/assets/js/ag-grid-community/**/*.js",
            "node_modules/angular-mocks/angular-mocks.js",
            "app/dist/public/resources/assets/js/config.js",
            "app/dist/public/resources/assets/js/app.js",
            "app/src/**/*.spec.js",
        ],
        frameworks: ["jasmine"],
        plugins: [
            require("karma-jasmine"),
            require("karma-chrome-launcher"),
            require("karma-jasmine-html-reporter"),
            require("karma-coverage"),
        ],
        client: {
            clearContext: false // leave Jasmine Spec Runner output visible in browser
        },
        coverageReporter: {
            dir: require("path").join(__dirname, "./coverage/traffic-portal"),
            reports: [{type: "html"}, {type: "text-summary"}],
            subdir: ".",
        },
        reporters: ["progress", "kjhtml"],
        port: 9876,
        colors: true,
        logLevel: config.LOG_INFO,
        autoWatch: true,
        browsers: ["Chrome"],
        singleRun: false,
        restartOnFileChange: true,
        customLaunchers: {
            ChromeHeadlessCustom: {
                base: "ChromeHeadless",
                flags: ["--no-sandbox", "--disable-gpu"],
            },
        },
    });
};
