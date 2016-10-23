# www.princebot.com
Source for www.princebot.com

Deployment scripts live in the <tt>scripts</tt> directory.

## Quickstart
```bash
git clone https://github.com/princebot/www.princebot.com
cd www.princebot.com

# Build docker image.
docker build -t www.princebot.com .

# Serve www.princebot.com
sudo docker run -d           \
    --name site              \
    --publish 80:8080        \
    --restart=unless-stopped \
    www.princebot.com serve
```

The scripts

## TODO
* Add SSL
* Add tests
* Add docs
