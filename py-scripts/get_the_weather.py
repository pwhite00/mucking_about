#!/usr/bin/env python
#
# query a weather service api and do stuff with the data.
#
#
#  openweathermap.org
#  apikey: f7f38c11ea8f2726b52b02f9d64dce3b
##########################################################
import requests


# lookup 5 day weather by zipcode
# exmaples api curl: api.openweathermap.org/data/2.5/forecast?zip={zip code},{country code}&APPID={key}
# http://api.openweathermap.org/data/2.5/forecast?zip=98075,us&APPID=f7f38c11ea8f2726b52b02f9d64dce3b
# JSON by default , &mode=xml or mode=html available as well
# Units is & units=standard, imperial, metric  - standard is default




