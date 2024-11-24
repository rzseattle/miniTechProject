import os
from ultralytics import YOLO


model = YOLO("yolo11n.pt")

file = "train/Apple Golden 1/0_100.jpg"

# Run inference on the image
results = model.predict(file)

# Print the predictions
print(results.pred)

# If you want to print the labels and confidence scores
# for result in results:
#     print(f"Label: {result['label']}, Confidence: {result['confidence']}")

