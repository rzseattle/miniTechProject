import os
from ultralytics import YOLO
import json
import pprint


model = YOLO("best.pt")

file = '/home/arturt/Downloads/avocado.jpg'
file = 'val/Salak/279_100.jpg'

# Run inference on the image
results = model.predict(file)

# Print the predictions
results = model.predict(file, verbose=False)

# Pretty print the predictions



predictions = []
for i, pred in enumerate(results[0].probs.top5):
    prediction = {
        "label": model.names[pred],
        "confidence": round(results[0].probs.top5conf[i].item(), 5)
    }
    predictions.append(prediction)

pprint.pprint(predictions)