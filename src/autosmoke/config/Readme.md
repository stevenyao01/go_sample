profinet need ../../../mqtt.conf ../../../device.sk ../../../server.crt
opcda need ../../../mqtt.conf ../../../device.sk


1, fins: 
    server, 10.111.111.114
    port,   9600
    dp,     CIO(BOOLEAN)

2, mewtocol:
    server, 10.111.111.111
    port,   6005 (6001-6005)
    dp,     R0(BOOLEAN)

3, bacnet:
    server,     10.111.71.27
    port,       47808 
    dp(import form liuyu):

4, modbus:
    server,         10.111.71.27
    port,           502
    dp(slaveid,1),  400001(dword), 400002(dword)

5, opcua:
    server, opc.tcp://10.111.71.27:48010
    dp,     Demo.Dynamic.Scalar.Boolean(boolean), namespace: 2

???6,opcda:
    server,     10.111.69.145
    user,       pc   
    password,   pwd
    clsid,      F8582CF2-88FB-11D0-B850-00C0F0104305
    dp,         Random.Int4(DINT)

7, profinet:
    server, 10.111.71.201
    port,   102
    dp,     M100(BOOLEAN)     

8, mc:
    server,         10.111.111.112
    port,           6000
    communication,  ascii
    net,            0
    netstation,     0
    connectId,      1
    dp,             M100(BOOLEAN)

9, iec104:
    server, 10.111.69.145
    port,   2404
    dp,     M_ME_NA(DWORD) 3 / M_DP_NA(CHAR) 2 / M_SP_NA(BOOLEAN) 1


10, filebeat:
    path,   ./a.log
