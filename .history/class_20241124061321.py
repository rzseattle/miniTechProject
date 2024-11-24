import os
from ultralytics import YOLO


model = YOLO("yolov8n-cls.yaml", task='classify') 


results = model.train(data="./", epochs=4)


model.export(format="onnx", file="yolov8n-cls.onnx")  