import os
from ultralytics import YOLO


model = YOLO("yolov8n-cls.yaml", task='classify') 


results = model.train(data="./", epochs=2)


model.export(format="tf")  