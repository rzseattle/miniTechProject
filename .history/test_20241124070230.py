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

for pred in results[0].probs:  # Iterate over predictions
    print(pred)
# for i in results[0]:
#     print(i)



