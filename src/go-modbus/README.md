# go-modbus

## About

This [Go](http://golang.org/) package provides [Modbus](http://en.wikipedia.org/wiki/Modbus) access for client (master) applications to communicate with server (slave) devices, over both [TCP/IP](http://www.modbus.org/docs/Modbus_Messaging_Implementation_Guide_V1_0b.pdf) and [Serial Line/RTU/ASCII](http://www.modbus.org/docs/Modbus_over_serial_line_V1_02.pdf) frame protocols.

Note that in modbus terminology, _client_ refers to the __master__ application or device, and the _server_ is the __slave__ waiting to respond to instructions, as shown in this transaction diagram:

![Modbus Transaction](http://i.imgur.com/Vgsqrb2.png)

This code was originally forked from [lubia/modbus](https://github.com/lubia/modbus) and repositioned as a pure client (master) library for use by controller applications.

### Installation and Usage

Install the package in your environment with these commands (the RTU code now depends on [goserial](https://github.com/tarm/goserial)):

```sh
go get github.com/tarm/goserial
go get github.com/dpapathanasiou/go-modbus
```

Next, build and run the examples:

 * [rtu-client.go](examples/rtu-client.go) for an RTU example
 * [ascii-client.go](examples/ascii-client.go) for an ASCII example
 * [tcp-client.go](examples/tcp-client.go) for a TCP/IP example


### Enabling the USB Serial Port adapter (RS-232) for RTU/ASCII Access

Slave devices which have [USB](http://en.wikipedia.org/wiki/Usb) ports for RTU access will not work immediately upon hot-plugging into a master computer.

For master devices running linux, the USB serial port adapter must be explicitly activated using the <tt>usbserial</tt> linux kernel module, as follows:

1. Immediately after plugging in the serial port USB, use <tt>dmesg</tt> to find the vendor and product ID numbers:

```
$ sudo dmesg | tail
```

There should be a line which looks like this:
	
```
[  556.572417] usb 3-1: New USB device found, idVendor=04d8, idProduct=000c
```

2. Use the <tt>usbserial</tt> linux kernel module to enable it, using the same vendor and product ID numbers from the dmesg output:

```
$ sudo modprobe usbserial vendor=0x04d8 product=0x000c
```

3. Confirm that the serial port is attached to a specific tty device file:

```
$ sudo dmesg | tail
```

There should now be a line like this:
   
```
[ 2134.866724] usb 3-1: generic converter now attached to ttyUSB0
```

which means that the serial port is now programmatically accessible via <tt>/dev/ttyUSB0</tt>

## References
- [Modbus Technical Specifications](http://www.modbus.org/specs.php)
- [Modbus Interface Tutorial](http://www.lammertbies.nl/comm/info/modbus.html)
- [Modbus TCP/IP Overview](http://www.rtaautomation.com/modbustcp/)
- [Modbus RTU Protocol Overview](http://www.rtaautomation.com/modbusrtu/)
- [Modbus ASCII Protocol Overview](http://www.simplymodbus.ca/ASCII.htm)

## Acknowledgements
- [Lubia Yang](http://www.lubia.me) for the [original modbus code](https://github.com/lubia/modbus) in Go
- [l.lefebvre](http://source.perl.free.fr/) for his excellent [modbus client](https://github.com/sourceperl/MBclient) and [server (slave device simulator)](https://github.com/sourceperl/mbserverd) code repositories
- [Tarmigan Casebolt](https://github.com/tarm/) for his [goserial](https://github.com/tarm/goserial) library, which resolved connection issues in RTU mode
- [modbusdriver.com](http://www.modbusdriver.com/) for their free [Diagslave Modbus Slave Simulator](http://www.modbusdriver.com/diagslave.html) tool
- [Mohammad Hafiz (mypapit)](https://plus.google.com/113437861006502895279?rel=author) for his well-written [How to enable USB-Serial Port adapter (RS-232) in Ubuntu Linux](http://blog.mypapit.net/2008/05/how-to-use-usb-serial-port-converter-in-ubuntu.html) blog post
