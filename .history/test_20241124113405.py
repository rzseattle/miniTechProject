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
results = model.predict(file, verbose=False)

# Pretty print the predictions

print(results[0].boxes)

predictions = []
for i, pred in enumerate(results[0].probs.top5):
    prediction = {
        "label": model.names[pred],
        "confidence": round(results[0].probs.top5conf[i].item(), 5)
    }
    predictions.append(prediction)

pprint.pprint(predictions)