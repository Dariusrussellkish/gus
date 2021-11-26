# Description: Calculate the average time it takes to serialize a message, send it to another machine,
# have that machine deserialize the message, and send an acknowledgement back.
# Usage: python calculate_statistics.py <file_name_of_dat_file>

import sys

with open(sys.argv[1], 'r') as f:
    content = f.read().splitlines()
    times = [int(time) for time in content]
    requests = len(times)
    average = sum(times) / requests
    print('average time in microseconds for {} requests: {}'.format(requests, average))
