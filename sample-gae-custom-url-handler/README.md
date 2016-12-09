# sample-gae-custom-url-handler
Sample Google Appengine App that will handle Sertone Data Delivery via Custom URL method

# Installation
git clone https://github.com/sertone/sample-gae-custom-url-handler.git
cd sample-gae-custom-url-handler

# Create an appengine project id
See Google App Engine Go Standard Environment Documentation
https://cloud.google.com/appengine/docs/go/

# Update app.yaml
application: your-gae-project-id

# Execute appcfg.py
cd sample-gae-custom-url-handler
appcfg.py --application=your-gae-project-id --email=<google account> --oauth2 update .

# Custom URL
https://your-gae-project-id.appspot.com/sertone
