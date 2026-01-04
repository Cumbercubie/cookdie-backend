import json
def read_first_n_rows(path: str, n: int = 1, print_data=True):
    print(f"--- Reading: {path} ---")
    records = []
    with open(path, "r") as f:
        for i, line in enumerate(f):
            if i >= n:
                break
            records.append(json.loads(line))
            if print_data:
                data = json.loads(line)
                print(data)
    return records
def main():
    # read_first_n_rows("../../data/yelp_dataset/yelp_academic_dataset_business.json")
    data = read_first_n_rows("../../data/yelp_dataset/yelp_academic_dataset_checkin.json")
    # data = read_first_n_rows("../../data/yelp_dataset/yelp_academic_dataset_review.json")
    # read_first_n_rows("../../data/yelp_dataset/yelp_academic_dataset_tip.json", 5)
    # data = read_first_n_rows("../../data/yelp_dataset/yelp_academic_dataset_user.json", 1, False)
    with open("checking.json", "w") as f:
        f.write(json.dumps(data))


if __name__ == "__main__":
    main()
