class ItemModel {
  String? id;
  String? name;
  String? creationDate;
  String? quantity;
  String? expirationDate;

  ItemModel({
    this.id,
    this.name,
    this.creationDate,
    this.quantity,
    this.expirationDate,
  });

  ItemModel.fromJson(Map<String, dynamic> json) {
    id = json['id'];
    name = json['name'];
    creationDate = json['creation_date'];
    quantity = json['quantity'];
    expirationDate = json['expiration_date'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['id'] = id;
    data['name'] = name;
    data['creation_date'] = creationDate;
    data['quantity'] = quantity;
    data['expiration_date'] = expirationDate;
    return data;
  }
}
