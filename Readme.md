# Memory Alert

[👉 Versão em português aqui](Readme.pt.md)

This is a simple utility written in Go to monitor the amount of available memory on a Linux system. When the available memory falls below a certain threshold, a notification is sent.

## Features
Monitors the available memory on the system.
Sends an alert notification when the available memory is less than 1500 MB.

Uses the notify package to display notifications on the desktop.
Prerequisites

This program is designed for Linux systems, as it relies on the /proc/meminfo file to obtain memory information.
Go installed on the system. You can download and install Go from golang.org.

How to Use
Clone the repository or download the source code to your working directory:

```

git clone https://github.com/your_user/memory-alert.git
cd memory-alert

```
Make sure you have an image named ram.png in the same directory where the executable will be run. This image will be used in the notification.

Compile the program using the command:

```
go build -o memory-alert
```

Run the program:

```
./memory-alert
```

Set memory-alert to start when your window manager starts

## Customization
The program will continuously monitor the available memory every 2 seconds and will display a notification when the available memory is less than 1500 MB.

Memory Limit: You can change the memory limit by passing the value via the command line.

Check Interval: The interval between memory checks is currently 2 seconds. You can adjust this by changing the value of time.Sleep(2 * time.Second) to a different value.

## License
This project is licensed under the MIT License - see the LICENSE file for more details.