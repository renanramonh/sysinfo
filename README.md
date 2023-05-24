[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)]()
[![Generic badge](https://img.shields.io/badge/Status-Alpha-red.svg)]()
[![GPLv3 license](https://img.shields.io/badge/License-Apache-purple.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://github.com/renanramonh)


# sysinfo

The jsonsort is a command-line tool written in Go that allows you to visualize system stats


## features
- returns CPU usage for each core
- returns storage usage
- returns amount of RAM

## Installation

To install `sysinfo`, follow these steps:

1. Make sure you have Go installed on your system. You can download and install it from the official website: https://golang.org/dl

2. Run the following command to install the JSON Sorter CLI:
```bash
go install github.com/renanramonh/sysinfo
```

3. The executable binary will be placed in your Go binary directory (`$GOPATH/bin`) after this command downloads the source code and compiles it..

4. (Optional) To facilitate access, add the Go binary directory to your system's PATH environment variable.
   Edit your shell configuration file (`~/.bashrc`, `~/.bash_profile`, or `~/.zshrc`) and add the following command: ``export PATH=$PATH:$GOPATH/bin``.
   Save the file and restart your terminal or run source `~/.bashrc` (or the corresponding command for your shell) to apply the changes.

## Usage
To run sysinfo use the following command:
```bash
sysinfo
```

the output will look like that
```txt
~$ sysinfo
System Information:
- Memory: 11.37 MB bytes
- Storage Size: 466.23 GB
- Storage Free: 219.51 GB
- Number of CPUs: 8
- CPU Usage Core 0: 21.43%
- CPU Usage Core 1: 39.39%
- CPU Usage Core 2: 13.68%
- CPU Usage Core 3: 46.94%
- CPU Usage Core 4: 17.53%
- CPU Usage Core 5: 15.05%
- CPU Usage Core 6: 35.64%
- CPU Usage Core 7: 18.37%

```

## Future Improvements
- Memory usage
- Export as json
