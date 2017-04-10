# Go  
  
	Some of my projects developed with Golang  
	I use this repository to track my growth  
  
### Colours  
**colours.go**  
####	Necessary setup  
    No additional packages required  
####	Short Workflow Description  
	Accepts an image as parameter (png or jpeg), and gives back the most frequent colours  
####	Short Algorithm Description  
    Executes a pixel by pixel scan of the image, uses a hash function to speed up the process  
  
### GMail Api Integration with Mail Sending Example  
**mailsender.go**  
####	Necessary setup  
    Follow [this] Google quickstart to setup your mail  
####	Short Workflow Description  
	Uses the specified (sender) account to send a mail to the specified (receiver) account. Needs to be authorized!  
####	Short Algorithm Description  
    Search for client_secret.json file, request a token to the google mail api and then send the mail  
    
    
    
    
[this]: http://google.com