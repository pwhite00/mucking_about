#!/usr/bin/env python
#
#  -- Collect facts around a system using python and not calling Facter
#
#  -- general familiarizzation with python and *nix system commands
#
#
#
###############################################################################
# collect the following:
# hostname
# OS family , majorVersion, MinorVersion
# uptime
# memory : total, free, swap
# disks: list , usage
# health: ?
# network interaces inventory
# network interace usage
# processor: id , usage
# load
# iostat
#
# write to terminal
# write to logger
# call inidividual tests or all.
# import commands -- don't use commands it's deprecated. use subprocess
# create a dictionary for all facts
# create subdictinoanires for individual fact groups
# output as json

import subprocess
import optparse
import json




def gather_id_facts():
    system_id = {}
    # reset test to true
    system_id['test'] = True

    # get hostname
    system_id['hostname'] = subprocess.Popen(['uname', '-n'], stdout=subprocess.PIPE)

    print system_id
    #return system_id


def gather_os_facts():
    os = {}
    # set test to true
    os['test'] = True

    # determine os family
    os['family'] = subprocess.Popen(['uname', '-s'], stdout=subprocess.PIPE)

    if os['family'] == 'Darwin':
        os['osName'] = subprocess.Popen(['sw_vers', '-productName'], stdout=subprocess.PIPE)
        os['osMajorRev'] = subprocess.Popen(['sw_vers', '-productVersion'], stdout=subprocess.PIPE)
        os['osMinorRev'] = subprocess.Popen(['sw_vers', '-buildVersion'], stdout=subprocess.PIPE)
    elif os['family'] == 'Linux':
        o['osName'] = subprocess.Popen(['lsb_release', '-i'], stdout=subprocess.PIPE)
        os['osMajorRev'] = subprocess.Popen("lsb_release -r | awk {'print $1'} | cut '.' -f 1", shell=True, stdout=subprocess.PIPE)
        os['osMinorRev'] = subprocess.Popen("lsb_release -r | awk {'print $1'} | cut '.' -f 2", shell=True, stdout=subprocess.PIPE)
    else:
        os['osName'] = 'not_collected.'
        os['osMajorRev'] = 'not_collected.'
        os['osMinorRev'] = 'not_collected.'

    return os


def gather_cpu_facts():
    cpu = {}
    # gather some such cpu facts and stuff
    # set test to true
    cpu['test'] = True

    return cpu


def gather_memory_facts():
    memory = {}
    # gather some such memory facts and stuff
    # set test to true
    memory['test'] = True

    return memory


def gather_storage_facts():
    storage = {}
    # gather all that handy disk info and stuff
    # set test to true
    storage['test'] = True

    return storage


def gather_network_facts():
    network = {}
    # also gather some handdy network info... and stuff
    # set test to true
    network['test'] = True

    return network


def collect_all_facts(system_id,):
    all_facts = {}
    # check each dictionary test and if True pull into all_facts.
    #id
    if system_id['test'] == True:
        all_facts['os_facts'] = system_id

    #os
    if os['test'] == True:
        all_facts['os_facts'] = os

    #cpu
    if cpu['test'] == True:
        all_facts['os_facts'] = cpu

    #memory
    if memory['test'] == True:
        all_facts['os_facts'] = memory

    #storage
    if storage['test'] == True:
        all_facts['os_facts'] = storage

    #network
    if network['test'] == True:
        all_facts['os_facts'] = network


    # start doing stuff with the data
        print all_facts

gather_id_facts
#gather_os_facts
#gather_cpu_facts
#gather_memory_facts
#gather_storage_facts
#gather_network_facts
#collect_all_facts()




