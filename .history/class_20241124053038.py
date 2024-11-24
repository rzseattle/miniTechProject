import os
from ultralytics import YOLO


model = YOLO("yolov8n-cls.pt")


results = model.train(data="./dataset", epochs=2)


model.export()  # export to TensorFlow