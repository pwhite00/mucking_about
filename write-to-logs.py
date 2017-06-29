#!/usr/bin/env python
#
# write a logger that can choose between syslog and a locale log file.
#
#
###############################################################################
import logging
import logging.handlers
import syslog
from optparse import OptionParser
from time import gmtime, strftime


parser = OptionParser()
parser.add_option('-l', '--local-log', action='store', type='string', dest='locallog',
                  help='Send Message to local log file.[ Include logging message following flag.]')
parser.add_option('-s', '--syslogger', action='store', type='string', dest='syslog',
                  help='Send Message to syslog. [ Include Logging message following flag.]')
parser.add_option('-b', '--both-logs', action='store', type='string', dest='bothlogs',
                  help='Write message to both local log file and syslog. [ Include logging message following flag.]')
(options, args) = parser.parse_args()

locallog = 'local.log'
time_stamp = strftime("%Y-%m-%d %H:%m:%s", gmtime())
log_level = 'logging.DEBUG'


def write_to_local_log(message, logname, time_stamp):
    # take the message and write it to a local log defined above.
    logging.basicConfig(filename=logname, level=logging.DEBUG)
    logmessage = time_stamp + message
    logging.info(logmessage)

def write_to_syslog(message):
    # take the message and write to syslog.
    syslog.openlog(logoption=syslog.LOG_PID, facility=syslog.LOG_SYSLOG)
    syslog.syslog(syslog.LOG_INFO, message)

def write_to_all_logs(message, logname, time_stamp):
    # take the message and logname and write to both local and syslog
    write_to_syslog(message)
    write_to_local_log(message, logname, time_stamp)

def write_to_syslog_using_logging(message):
    # now without using the syslog libray write to syslog using logging and logging.handlers
    

# route input to correct function.
if options.locallog:
    write_to_local_log(options.locallog, locallog, time_stamp)
elif options.syslog:
    write_to_syslog(options.syslog)
elif options.bothlogs:
    write_to_all_logs(options.bothlogs, locallog, time_stamp)

