import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class ListStatistics extends StatefulWidget {
  const ListStatistics({super.key});

  @override
  State<ListStatistics> createState() => _ListStatisticsState();
}

class _ListStatisticsState extends State<ListStatistics> {
  late Future<ResponseData> futureData;

  @override
  void initState() {
    super.initState();
    futureData = fetchStatistics().then((data) => ResponseData.fromJson(data));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Quantidade Minima'),
      ),
      body: FutureBuilder(
        future: futureData,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            return ListView.builder(
              itemCount: snapshot.data!.productList.length,
              itemBuilder: (context, index) {
                return Card(
                  elevation: 2,
                  color: Colors.grey[300],
                  child: ListTile(
                    title: Center(
                      child: Text(snapshot.data!.productList[index].name),
                    ),
                    // subtitle: Text(snapshot.data!.minimalQuantity.toString()),
                  ),
                );
              },
            );
          } else if (snapshot.hasError) {
            return Center(child: Text('${snapshot.error}'));
          }
          return const Center(child: CircularProgressIndicator());
        },
      ),
    );
  }
}

Future<Map<String, dynamic>> fetchStatistics() async {
  final response = await http
      .get(Uri.parse('https://despensa.onrender.com/api/stock/statistics'));

  if (response.statusCode == 200) {
    return json.decode(response.body);
  } else {
    throw Exception('Failed to load data');
  }
}

class Item {
  final String id;
  final String name;

  Item({required this.id, required this.name});

  factory Item.fromJson(Map<String, dynamic> json) {
    return Item(
      id: json['id'],
      name: json['name'],
    );
  }
}

class ResponseData {
  final int minimalQuantity;
  final int affectedProducts;
  final List<Item> productList;

  ResponseData({
    required this.minimalQuantity,
    required this.affectedProducts,
    required this.productList,
  });

  factory ResponseData.fromJson(Map<String, dynamic> json) {
    List<Item> products = [];
    for (var productJson in json['product_list']) {
      products.add(Item.fromJson(productJson));
    }

    return ResponseData(
      minimalQuantity: json['minimal_quantity'],
      affectedProducts: json['affected_products'],
      productList: products,
    );
  }
}
