import os
from ultralytics import YOLO


model = YOLO("yolov8n-cls.pt")

file = "train/Apple Golden 1/0_100.jpg"

# Run inference on the image
results = model.predictor(file)

# Print the predictions
print(results)