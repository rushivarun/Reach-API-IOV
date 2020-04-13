import subprocess
import time
import argparse

parser = argparse.ArgumentParser(description='Display WLAN signal strength.')
parser.add_argument(dest='interface', nargs='?', default='wlp2s0',
                    help='wlan interface (default: wlp2s0)')
args = parser.parse_args()

print('\n---Press CTRL+Z or CTRL+C to stop.---\n')

while True:
    cmd = subprocess.Popen('sudo iw wlp2s0 scan', shell=True,
                           stdout=subprocess.PIPE)
    for line in cmd.stdout:
        print(line)
        if b'signal' in line:
            print(line.lstrip(b' '),)
        elif b'Not-Associated' in line:
            print('No signal')
    time.sleep(1)