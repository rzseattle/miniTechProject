import os
from ultralytics import YOLO
import json
import pprint
import cv2


model = YOLO("best.pt")

file = '/home/arturt/Downloads/avocado.jpg'
file = 'val/Salak/279_100.jpg'
file = 'Fruit-Images-Dataset/test-multiple_fruits/apple_apricot_peach_peach(flat)_pomegranate_pear_plum_2.jpg'

# Run inference on the image
results = model.predict(file)

# Print the predictions
# Draw bounding boxes on the image

image = cv2.imread(file)
for box in results[0].boxes:
    x1, y1, x2, y2 = map(int, box.xyxy[0])
    label = model.names[box.cls]
    confidence = box.conf.item()
    
    # Draw the bounding box
    cv2.rectangle(image, (x1, y1), (x2, y2), (0, 255, 0), 2)
    
    # Put the label and confidence on the box
    cv2.putText(image, f'{label} {confidence:.2f}', (x1, y1 - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.9, (0, 255, 0), 2)

# Save the image with bounding boxes
cv2.imwrite('/home/arturt/projects/miniTech/output.jpg', image)
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