import os
from ultralytics import YOLO


model = YOLO("best.pt")

file = '/home/arturt/Downloads/avocado.jpg'
file = 'val/Salak/279_100.jpg'

# Run inference on the image
results = model.predict(file)

# Print the predictions
print(results.Probs)

# If you want to print the labels and confidence scores
# for result in results:
#     print(f"Label: {result['label']}, Confidence: {result['confidence']}")
