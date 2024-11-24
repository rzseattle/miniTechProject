import os
from ultralytics import YOLO


model = YOLO("yolov8n-cls.pt")


//results = model.train(data="./", epochs=1)


model.export(format="tf")  # export to TensorFlow