Comp 429: Computer Network Software
Section 7815 – SP2022
Programming Assignment 1: A Chat Application
for Remote Message Exchange

Group members:
Vladimir Sterlin
Cristian Ramirez
---------------------------------------------------------
Contribution 

Vladimir began the layout of the program by setting up several
go files, including main, app, client, server, address, remove,
and trim. Then he started to fill out the functions and 
requirements in each file to make the app work. Cristian was tasked
with setting up the functionality such as the user interface
and the several commands the user can input that was mentioned 
in the instructions, such as help, myip, connect, list, and many
others. He also dealt with the error handling, including if there
was an error in the connection, warning the user they can't do self-connection
or use the same IP, and various others. 

In main.go, the file imports the methods and classes necessary for the
program. Vladimir decided to make the main method to create the 
app and run it there as well. 

In the app.go file, this is where we house the main code of the programs.
There is a struct called, app that stores the client and server, and the
connections that the program has with other devices via network, which
inlcudes their IP address and port number. Vladimir made the layout, including 
having the NewApp method to establish the port number to be used
for connection and creating an array to store all connections this socket has.
He also wrote the Listen call for the server and having the switch case for the 
several actions that can be done, based on the user input while Cristian 
was adding the user interface and making it more user friendly. The user is shown
a list of commands they can type to perform an action. While Vladimir was typing the
methods that would be used for each case, Cristian would type the appropriate 
display message and any error handling it needs to have. One example is
using an error handling when the user types in "terminate" without using "list" first
since the user would not know which connection they want to close. Another example
would be Vladimir writing the methods to establish a connection while Cristian
makes sure there are no self-connection or duplicate connections.
There was several helper methods that Vladimir has created to help out with the app method,
such as listing the current connections the socket is associated, reading the user input and
making it easier for the program to read by removing spaces or new lines, closing all 
connections when the user exits and terminating a specific connection.  

In client.go, Vladimir wrote three methods that are meant for the client,
including connecting the client with the server, able to send messages to 
the server, and closing the connection when it is necssary, such as
terminating a specific connection. There is also a struct for the client
only focuses on the connections the client has. Cristian typed the error 
handling for some actions, such as if the program was not able to send 
a message or there was an error trying to connect. We also decided 
to display a message to the user when a connection has been established. 

In the server.go, we have also created a struct for server, in which it contains
the connections, IP address and port number for each server it wants to connect.
Such function includes the Listen, handleClient to deal with the messages it received. 
Similar fashion where most of the code was written by Vladimir and the error handling was dealt
by Cristian. 

When it came to the last three go files, which are considered the utility files,
they were created to help out with certain actions, such as obtaining the user's IP,
building the address using the port and IP address, properly removing a connection's IP address
from the list when the user uses "terminate" so the list's size is reduce by one. And the final 
method is trimming strings to be easier to deal with in the main app, such as removing the new lines.

---------------------------------------------------------
There are two ways to run the file, one that uses the Makefile
and one that doesn't.


Option 1: Using the Makefile

If you prefer to use the Makefile, then the prerequisites is to 
install the GO compiler, and the Make software. Once those are 
complete, the user can build the file using the Make command. Once
the file is built, the user can run the file using the make command and
the name of the file. The user can also use ./chat #### where the #'s
represent the port number.

Option 2: Using the codes instead

If you prefer to not use the Makefile, then you can use a 
source code editor such as Visual Studios. Once that has
been installed, you will download several extensions that
supports GO language, including Code Runner and Go language
for VS. You will also need to install the Go compiler, which
is found here: https://go.dev/dl/

Once the necessary resources are installed, you can open up
Visual Studios and use the interal terminal to run the code.
You will access the directory TCP-Chat-App and once you
reached to this location, run the code using "go run
main.go ####" where the #'s is the port number you want 
to use. 

NOTE: It is also possible to use the command prompt instead
of Visual Studio if you desire, as long as you have the 
GO Compiler installed and the path way directed to the 
cmd, which should be done automatically. It will run
using the same command line. 

Command Line: go run main.go ####