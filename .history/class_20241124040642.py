import os
from ultralytics import YOLO


model = YOLO("yolov8n-cls.pt", device="cuda")


