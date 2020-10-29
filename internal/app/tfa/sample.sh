#!/bin/bash
npm install speakeasy qrcode > "/dev/null" 2>&1
echo $(node internal/app/tfa/2fa.js)
