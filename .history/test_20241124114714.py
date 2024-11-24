import os
from ultralytics import YOLO
import json
import pprint


model = YOLO("best.pt")

file = '/home/arturt/Downloads/avocado.jpg'
file = 'val/Salak/279_100.jpg'
file = 'Fruit-Images-Dataset/test-multiple_fruits/apple_apricot_peach_peach(flat)_pomegranate_pear_plum_2.jpg'

# Run inference on the image
results = model.predict(file)

# Print the predictions
# Draw bounding boxes on the image

# Extract bounding boxes, classes, names, and confidences
boxes = results[0].boxes.xyxy.tolist()
classes = results[0].boxes.cls.tolist()
names = results[0].names
confidences = results[0].boxes.conf.tolist()

# Iterate through the results
for box, cls, conf in zip(boxes, classes, confidences):
    x1, y1, x2, y2 = box
    confidence = conf
    detected_class = cls
    name = names[int(cls)]

# predictions = []
# for i, pred in enumerate(results[0].probs.top5):
#     prediction = {
#         "label": model.names[pred],
#         "confidence": round(results[0].probs.top5conf[i].item(), 5)
#     }
#     predictions.append(prediction)

# pprint.pprint(predictions)