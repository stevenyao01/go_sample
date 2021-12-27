#!/bin/bash
# chkconfig:   2345 90 10
# description:  agent


agent=""
pwd1=""
filePath="/etc/init.d/agent"
agentName="agent"


function help()
{
    echo "Usage: $0 [parameter]"
    echo ""
    echo "Example:"
    echo "      "$0 "install"
    echo ""
    echo "Parameter:"
    echo "	install"
}

function serviceHelp()
{
    echo "Usage service agent [parameter]"
    echo ""
    echo "Example:"
    echo "	service agent start"
    echo "Parameter:"
    echo "	start,stop,restart,status,uninstall"
}

function initAgent()
{
    files=$(ls $folder)
    for file in $files
    do
        if [[ $file =~ "EdgeAgent_" && ! $file =~ ".zip" ]];
        then
            if [ ! -x $file ];
            then
                chmod +x $file
            fi
	    sed -i "6c agent=$file" $filePath
	    sed -i "6c agent=$file" agent.sh
        fi
    done
    pwd1=$(cd "$(dirname "$0")";pwd)
    sed -i "7c pwd1=$pwd1" $filePath
    sed -i "7c pwd1=$pwd1" agent.sh
}

function installAgent()
{
    if [ ! -f "$filePath" ];
    then
	#rm $filePath
	cp agent.sh $filePath
	initAgent
	centos=$(echo   "centos"   |   tr   [a-z]   [A-Z])
	ubuntu=$(echo   "ubuntu"   |   tr   [a-z]   [A-Z])
	system=`cat /etc/*release`
	sys=$(echo   $system | tr [a-z] [A-Z])
	if [[ $sys =~ $centos ]];
	then
            echo "centos"
	    centosConfig
	elif [[ $sys =~ $ubuntu ]];
	then
            echo "ubuntu"
	    ubuntuConfig
	else
            echo "other"
	    centosConfig
	fi
	source /etc/profile
	service $agentName start
    else
	echo "Bootstrap service already exists..."
	exit
    fi
}

function centosConfig()
{
    cd /etc/init.d/
    chmod +x $agentName
    chkconfig --add $agentName
    chkconfig $agentName on
}

function ubuntuConfig()
{
    cd /etc/init.d/
    chmod +x $agentName
    update-rc.d $agentName defaults 50
}

function uninstallAgent()
{
    if [ -f "$filePath" ];
    then
	cd /etc/init.d/
        service $agentName stop
	centos=$(echo   "centos"   |   tr   [a-z]   [A-Z])
        ubuntu=$(echo   "ubuntu"   |   tr   [a-z]   [A-Z])
        system=`cat /etc/*release`
        sys=$(echo   $system | tr [a-z] [A-Z])
        if [[ $sys =~ $centos ]];
        then
            echo "centos $agentName"
            chkconfig $agentName off
        elif [[ $sys =~ $ubuntu ]];
        then
            echo "ubuntu"
	    update-rc.d -f $agentName remove
        else
            echo "other"
	    chkconfig $agentName off
        fi
	rm $agentName	
    else
        echo "Bootstrap service already exists..."
	exit
    fi
}

function startAgent()
{
    pid=`ps -ef |grep $pwd1/.$agent|grep -v grep |awk '{print $2,$3}'`
    if [[ $pid == "" ]];
    then
	cd $pwd1
	nohup ./$agent >/dev/null 2>&1 &
    else
	echo "Service started pid : $pid"
    fi
}

function stopAgent()
{
    pid=`ps -ef |grep $pwd1/.$agent|grep -v grep |awk '{print $2,$3}'`
    if [[ $pid != "" ]];
    then
	kill -9 $pid
    else
	echo "No shutdown process..."    
    fi
}
function statusAgent()
{
    #ps -ef | grep $agent |grep -v grep
    pid=`ps -ef |grep $pwd1/.$agent|grep -v grep |awk '{print $2,$3}'`
    if [[ $pid == "" ]];
    then
	echo "agent not running "
    else
        echo "agent is running,pid : $pid"
    fi

}

# check args
if [ $# -eq 0 ];
then
    if [[ $1 -eq "help"  ||  $1 -eq "--help"  ||  $1 -eq "-h" ]];
    then
	if [ -f "$filePath" ];
	then
	    serviceHelp
	    exit
	else
	    help
	    exit
	fi
    else
	if [ ! -f "$filePath" ];
	then
            service $agentName start
            exit
        fi
    fi
fi

if [ $# -ne 1 ];
then
    help
    exit
else
    if [ -f "$filePath" ];
    then
    	if [ "$1" == "uninstall" ];
	then
            uninstallAgent
            exit
	elif [ "$1" == "start" ];
	then
            echo "start agent..."
            startAgent
            exit
	elif [ "$1" == "stop" ];
	then
            echo "stop agent..."
            stopAgent
            exit
	elif [ "$1" == "restart" ];
	then
            echo "restart agent..."
            stopAgent
            echo "stop success start agent..."
            startAgent
	elif [ "$1" == "status" ];
	then
	    echo "agent status..."
	    statusAgent
	else
            echo "The service is installed. Please use the following command "
            echo ""
            serviceHelp
            exit
	fi
    else
	if [ "$1" == "install" ];
        then
	    echo "start install..."
            installAgent
            exit
	else
	    help
	    exit
	fi
    fi
fi

