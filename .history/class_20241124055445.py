import os
from ultralytics import YOLO


model = YOLO("yolov8n.yaml", task='classify') 


results = model.train(data="./", epochs=2)


model.export()  # export to TensorFlow