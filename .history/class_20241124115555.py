import os
from ultralytics import YOLO


#model = YOLO("yolov8n-cls.yaml", task='classify') 
model = YOLO("yolov8n-cls.yaml", task='classify') 


results = model.train(data="./", epochs=25)


model.export()  