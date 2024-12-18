import os
from ultralytics import YOLO
import json
import pprint


model = YOLO("runs/classify/train2/weights/last.pt")
#model = YOLO("yolo11n.pt")


# Get the list of files in the "./test" folder
image_folder = "./test"
image_files = [os.path.join(image_folder, f) for f in os.listdir(image_folder)]
# print (image_files)
# print (len(image_files))

# file = '/home/arturt/Downloads/avocado.jpg'
# file2 = 'val/Salak/279_100.jpg'
# file3 = 'Fruit-Images-Dataset/test-multiple_fruits/apple_apricot_peach_peach(flat)_pomegranate_pear_plum_2.jpg'

# Run inference on the image
results = model.predict(image_files)

# Print the predictions
# Draw bounding boxes on the image

# Extract bounding boxes, classes, names, and confidences
# boxes = results[0].boxes.xyxy.tolist()
# classes = results[0].boxes.cls.tolist()
# names = results[0].names
# confidences = results[0].boxes.conf.tolist()

# # Iterate through the results
# for box, cls, conf in zip(boxes, classes, confidences):
#     x1, y1, x2, y2 = box
#     confidence = conf
#     detected_class = cls
#     name = names[int(cls)]

predictions = []
for result in results:
    resultData = {
        "file": result.path,
        "probs": []
    }
    
    for i, pred in enumerate(result.probs.top5):
        resultDataProb = {
            "label": model.names[pred],
            "confidence": round(results[0].probs.top5conf[i].item(), 5)
        }
        resultData["probs"].append(resultDataProb)

    predictions.append(resultData)

print(json.dumps(predictions, indent=4))