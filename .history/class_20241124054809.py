import os
from ultralytics import YOLO


model = YOLO("yolo11n.yaml")


results = model.train(data="./", epochs=2)


model.export()  # export to TensorFlow