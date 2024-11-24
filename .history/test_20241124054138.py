import os
from ultralytics import YOLO


model = YOLO("yolov8n-cls.pt")

file = "train/Apple Golden 1/0_100.jpg"

# Run inference on the image
results = model.predict(file)

# Print the predictions



# Define path to the image file
source = "path/to/image.jpg"

# Run inference on the source
results = model(file)  # list of Results objects
print(results)