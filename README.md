# Go_QR_Grid
For using the CLI arguments you first need to build the MainCLI.go file by using:

    go build .\MainCLI.go

after building the file you can use the following arguments to run the program

   "-BGsize int" For Size of Background Image Width and Height in Pixels
   
  "-QRnum int" For Number of QRcodes
  
  "-QRsize int" For Size of QRcodes Width and Height in Pixels
 
 Example of the full command should look something like this:
 
    .\MainCLI -BGsize 350 -QRnum 6 -QRsize 125
    
 An image called QRgrid.png will be made in the same folder as the program, that image is the final result
    
.

.

.

.

.
   
 For the GUI you need to build the mainGui.go file by using:
   
    go build .\maingui.go
    
 The other thing you need to have for the maingui.exe to run is the maingui.exe.manifest which you can also find in the project repository, just put the maingui.exe and the maingui.exe.manifest in the same folder and you are ready to go.
