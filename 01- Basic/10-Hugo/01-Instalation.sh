# Link https://gohugo.io/installation/linux/
# Debian 
sudo apt install hugo


# https://github.com/gohugoio/hugo
# Build the standard edition:
go install github.com/gohugoio/hugo@latest


# Build the extended edition:
# Turn on VPN from Iran
CGO_ENABLED=1 go install -tags extended github.com/gohugoio/hugo@latest
