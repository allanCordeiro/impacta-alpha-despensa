// class BalanceProducts {
// 	List<03052023>? l03052023;
// 	List<04052023>? l04052023;
// 	List<10052023>? l10052023;
// 	List<12052023>? l12052023;
// 	List<15052023>? l15052023;
// 	List<19052023>? l19052023;
// 	List<20042023>? l20042023;
// 	List<21042023>? l21042023;
// 	List<23052023>? l23052023;
// 	List<24042023>? l24042023;
// 	List<25042023>? l25042023;
// 	List<27042023>? l27042023;
// 	List<29052023>? l29052023;

// 	BalanceProducts({this.l03052023, this.l04052023, this.l10052023, this.l12052023, this.l15052023, this.l19052023, this.l20042023, this.l21042023, this.l23052023, this.l24042023, this.l25042023, this.l27042023, this.l29052023});

// 	BalanceProducts.fromJson(Map<String, dynamic> json) {
// 		if (json['03-05-2023'] != null) {
// 			l03052023 = <03052023>[];
// 			json['03-05-2023'].forEach((v) { l03052023!.add(new 03052023.fromJson(v)); });
// 		}
// 		if (json['04-05-2023'] != null) {
// 			l04052023 = <04052023>[];
// 			json['04-05-2023'].forEach((v) { l04052023!.add(new 04052023.fromJson(v)); });
// 		}
// 		if (json['10-05-2023'] != null) {
// 			l10052023 = <10052023>[];
// 			json['10-05-2023'].forEach((v) { l10052023!.add(new 10052023.fromJson(v)); });
// 		}
// 		if (json['12-05-2023'] != null) {
// 			l12052023 = <12052023>[];
// 			json['12-05-2023'].forEach((v) { l12052023!.add(new 12052023.fromJson(v)); });
// 		}
// 		if (json['15-05-2023'] != null) {
// 			l15052023 = <15052023>[];
// 			json['15-05-2023'].forEach((v) { l15052023!.add(new 15052023.fromJson(v)); });
// 		}
// 		if (json['19-05-2023'] != null) {
// 			l19052023 = <19052023>[];
// 			json['19-05-2023'].forEach((v) { l19052023!.add(new 19052023.fromJson(v)); });
// 		}
// 		if (json['20-04-2023'] != null) {
// 			l20042023 = <20042023>[];
// 			json['20-04-2023'].forEach((v) { l20042023!.add(new 20042023.fromJson(v)); });
// 		}
// 		if (json['21-04-2023'] != null) {
// 			l21042023 = <21042023>[];
// 			json['21-04-2023'].forEach((v) { l21042023!.add(new 21042023.fromJson(v)); });
// 		}
// 		if (json['23-05-2023'] != null) {
// 			l23052023 = <23052023>[];
// 			json['23-05-2023'].forEach((v) { l23052023!.add(new 23052023.fromJson(v)); });
// 		}
// 		if (json['24-04-2023'] != null) {
// 			l24042023 = <24042023>[];
// 			json['24-04-2023'].forEach((v) { l24042023!.add(new 24042023.fromJson(v)); });
// 		}
// 		if (json['25-04-2023'] != null) {
// 			l25042023 = <25042023>[];
// 			json['25-04-2023'].forEach((v) { l25042023!.add(new 25042023.fromJson(v)); });
// 		}
// 		if (json['27-04-2023'] != null) {
// 			l27042023 = <27042023>[];
// 			json['27-04-2023'].forEach((v) { l27042023!.add(new 27042023.fromJson(v)); });
// 		}
// 		if (json['29-05-2023'] != null) {
// 			l29052023 = <29052023>[];
// 			json['29-05-2023'].forEach((v) { l29052023!.add(new 29052023.fromJson(v)); });
// 		}
// 	}

// 	Map<String, dynamic> toJson() {
// 		final Map<String, dynamic> data = new Map<String, dynamic>();
// 		if (this.l03052023 != null) {
//       data['03-05-2023'] = this.l03052023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l04052023 != null) {
//       data['04-05-2023'] = this.l04052023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l10052023 != null) {
//       data['10-05-2023'] = this.l10052023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l12052023 != null) {
//       data['12-05-2023'] = this.l12052023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l15052023 != null) {
//       data['15-05-2023'] = this.l15052023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l19052023 != null) {
//       data['19-05-2023'] = this.l19052023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l20042023 != null) {
//       data['20-04-2023'] = this.l20042023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l21042023 != null) {
//       data['21-04-2023'] = this.l21042023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l23052023 != null) {
//       data['23-05-2023'] = this.l23052023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l24042023 != null) {
//       data['24-04-2023'] = this.l24042023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l25042023 != null) {
//       data['25-04-2023'] = this.l25042023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l27042023 != null) {
//       data['27-04-2023'] = this.l27042023!.map((v) => v.toJson()).toList();
//     }
// 		if (this.l29052023 != null) {
//       data['29-05-2023'] = this.l29052023!.map((v) => v.toJson()).toList();
//     }
// 		return data;
// 	}
// }

// class 03052023 {
// 	String? name;
// 	int? deductedQuantity;
// 	int? remainingQuantity;

// 	03052023({this.name, this.deductedQuantity, this.remainingQuantity});

// 	03052023.fromJson(Map<String, dynamic> json) {
// 		name = json['name'];
// 		deductedQuantity = json['deducted_quantity'];
// 		remainingQuantity = json['remaining_quantity'];
// 	}

// 	Map<String, dynamic> toJson() {
// 		final Map<String, dynamic> data = new Map<String, dynamic>();
// 		data['name'] = this.name;
// 		data['deducted_quantity'] = this.deductedQuantity;
// 		data['remaining_quantity'] = this.remainingQuantity;
// 		return data;
// 	}
// }