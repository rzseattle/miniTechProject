import os
from ultralytics import YOLO


#model = YOLO("yolov8n-cls.yaml", task='classify') 
model = YOLO("yolov8n-cls.pt", task='classify') 


results = model.train(data="./", epochs=4)


model.export()  