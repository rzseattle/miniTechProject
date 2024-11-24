import os
from ultralytics import YOLO
import json


model = YOLO("best.pt")

file = '/home/arturt/Downloads/avocado.jpg'
file = 'val/Salak/279_100.jpg'

# Run inference on the image
results = model.predict(file)

# Print the predictions
#print(results[0].names)

# Pretty print the predictions

for pred in results[0].boxes:  # Iterate over predictions
    class_id = int(pred.cls)  # Class ID
    confidence = float(pred.conf)  # Confidence score
    box = pred.xyxy  # Bounding box (x_min, y_min, x_max, y_max)

    print(f"Class: {model.names[class_id]}")  # Class name
    print(f"Confidence: {confidence:.2f}")
    print(f"Bounding Box: {box.numpy()}")
    print("-" * 40)

# for i in results[0]:
#     print(i)



