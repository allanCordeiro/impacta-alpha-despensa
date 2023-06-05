import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

import '../repository/implementations/http_api_repository.dart';

class BalanceItemPage extends StatefulWidget {
  final String? operationDate;
  final int? deductedQuantity;
  final int? remainingQuantity;
  final String id; // Novo parâmetro ID

  const BalanceItemPage({
    Key? key,
    this.operationDate,
    this.deductedQuantity,
    this.remainingQuantity,
    required this.id, // Novo parâmetro ID
  }) : super(key: key);

  @override
  State<BalanceItemPage> createState() => _BalanceItemPageState();
}

class _BalanceItemPageState extends State<BalanceItemPage> {
  final HttpApiRepository client = HttpApiRepository();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Historico'),
      ),
      body: FutureBuilder(
        future: getBalance(widget.id),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          } else if (snapshot.hasError) {
            return Center(
              child: Text('Error: ${snapshot.error}'),
            );
          } else {
            // Use the balance data to build the ListView
            List<dynamic> data = snapshot.data as List<dynamic>;
            return GridView.builder(
              primary: false,
              padding: const EdgeInsets.all(20),
              itemCount: data.length,
              gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 3,
                childAspectRatio: 1,
              ),
              itemBuilder: (context, index) {
                return Card(
                  elevation: 2,
                  color: Colors.grey[200],
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Row(
                        children: [
                          Expanded(
                              child: Text(
                                  data[index]['operation_date'].toString())),
                        ],
                      ),
                      Row(
                        children: [
                          Text(data[index]['deducted_quantity'].toString()),
                        ],
                      ),
                      Row(
                        children: [
                          Text(data[index]['reimaining_quantity'].toString()),
                        ],
                      ),
                    ],
                  ),
                );
              },
            );
          }
        },
      ),
    );
  }

  Future<List<dynamic>> getBalance(String id) async {
    final response = await http.get(
        Uri.parse('https://despensa.onrender.com/api/products/$id/balance'));

    if (response.statusCode == 200) {
      final decodedData = json.decode(response.body);
      return (decodedData as List<dynamic>);
    } else {
      throw Exception('Failed to fetch balance');
    }
  }
}
