import os
from ultralytics import YOLO


model = YOLO("best.pt")

file = '/home/arturt/Downloads/avocado.jpg'
file = 'val/Salak/279_100.jpg'

# Run inference on the image
results = model.predict(file)

# Print the predictions
print(results)

# Pretty print the predictions
for result in results:
    print(f"Class: {result['class']}, Confidence: {result['confidence']:.2f}, Bounding Box: {result['bbox']}")

