# Youtube MP3 Go Downloader
The Material Design Html5 Go Based Youtube MP3 Download
# What is this?
This is an open-source,Material Design,Golang-Based, Youtube to MP3 Converter/Downloader.
# Screenshots

# How does it work?
The golang server downloads and converts youtube videos into MP3 and serves the files.
Thanks to the html5/javascript frontend you can ask the golang-server to download files.
# Dependencies
youtube-dl ffmpeg/avconv golang
    sudo apt-get install youtube-dl ffmpeg
And a webserver (apache2,nginx or similar)
# How to install
-Clone my repo
    git clone http://github.com/nik012003/Youtube-MP3-Go && cd Youtube-MP3-Go/
-Copy Frontend files to your webserver
    cp Frontend/* /var/www/html
-Run main.go
    cd Backend-Server && go run main.go
-Enjoy
# Before deploying
You may need to change IPs and ports in main.go and index.html to work in your own setup.
# Contributing
Feel free to fork and push commits and Feature Requests to this project.
