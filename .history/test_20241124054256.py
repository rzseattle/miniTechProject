import os
from ultralytics import YOLO


model = YOLO("yolov8n-cls.pt")

file = "train/Apple Golden 1/0_100.jpg"

# Run inference on the image
results = model.predict(file)

# Print the predictions
print(results)

# If you want to print the labels and confidence scores
for result in results:
    print(f"Label: {result['label']}, Confidence: {result['confidence']}")

# Alternatively, you can use the model directly on the image
source = "path/to/image.jpg"
results = model(source)  # list of Results objects

# Print the results
print(results)

# If you want to print the labels and confidence scores from the results
for result in results:
    print(f"Label: {result['label']}, Confidence: {result['confidence']}")