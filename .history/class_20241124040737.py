import os
from ultralytics import YOLO


model = YOLO("yolov8n-cls.pt", device="cuda")


results = model.train(data="./", epochs=2)
