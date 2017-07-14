echo "*** Changing file permissions to user '$USER'"
sudo chown -R $USER:$USER ./
sudo chmod +x *

echo "*** Permissions changed"
