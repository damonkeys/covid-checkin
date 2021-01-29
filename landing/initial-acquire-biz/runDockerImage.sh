docker run -p 127.0.0.1:80:4444 ${{ secrets.REGISTRY_SERVER }}/chckr/landing-$(basename $PWD)
