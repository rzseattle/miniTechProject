pip install matplotlib numpy "tensorflow[and-cuda]" matplotlib --break-system-packages
pip install pandas seaborn  --break-system-packages
SKLEARN_ALLOW_DEPRECATED_SKLEARN_PACKAGE_INSTALL=True pip install sklearn --break-system-packages
#opencv-python
#keras-utils
# pip install matplotlib numpy tensorflow pandas seaborn sklearn --break-system-packages
pip3 install -U scikit-learn scipy matplotlib

#pip3 install nvidia-tensorrt
#python3 -m venv env
#source env/bin/activate
#pip3 install tensorflow



wget https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2404/x86_64/cuda-keyring_1.1-1_all.deb
sudo dpkg -i cuda-keyring_1.1-1_all.deb
sudo apt-get update
sudo apt-get -y install cudnn


cp -r Fruit-Images-Dataset/Training train
cp -r Fruit-Images-Dataset/Test test




pip3 install transformers datasets torch torchvision