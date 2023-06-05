class BalanceIten {
  String? operationDate;
  int? deductedQuantity;
  int? reimainingQuantity;

  BalanceIten(
      {this.operationDate, this.deductedQuantity, this.reimainingQuantity});

  BalanceIten.fromJson(Map<String, dynamic> json) {
    operationDate = json['operation_date'];
    deductedQuantity = json['deducted_quantity'];
    reimainingQuantity = json['reimaining_quantity'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['operation_date'] = operationDate;
    data['deducted_quantity'] = deductedQuantity;
    data['reimaining_quantity'] = reimainingQuantity;
    return data;
  }
}
