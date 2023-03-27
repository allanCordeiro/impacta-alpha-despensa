import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

import 'item_register_page.dart';

class ListItems extends StatefulWidget {
  const ListItems({Key? key}) : super(key: key);

  @override
  _ListItemsState createState() => _ListItemsState();
}

class _ListItemsState extends State<ListItems> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Lista de Itens'),
      ),
      body: const FutureListBuild(),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => const ItemRegister(
                entryDate: null,
                expirationDate: null,
                id: null,
                quantity: null,
              ),
            ),
          ).then((value) {
            if (value == true) {
              setState(() {});
            }
          });
        },
        child: const Icon(Icons.add),
      ),
    );
  }
}

class FutureListBuild extends StatelessWidget {
  const FutureListBuild({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<dynamic>(
      future: getList(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return GridView.builder(
            itemCount: snapshot.data!.length,
            gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: 2,
            ),
            itemBuilder: (BuildContext context, int index) {
              return Card(
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Padding(
                      padding: const EdgeInsets.all(8.0),
                      child: Text(
                        snapshot.data[index]['name'],
                        style: const TextStyle(
                          fontWeight: FontWeight.bold,
                          fontSize: 16,
                          fontFamily: 'Montserrat',
                        ),
                      ),
                    ),
                    
                    Padding(
                      padding: const EdgeInsets.all(8.0),
                      child: Text(
                        snapshot.data[index]['expiration_date'],
                        style: const TextStyle(
                          fontWeight: FontWeight.bold,
                          fontSize: 16,
                          fontFamily: 'Montserrat',
                        ),
                      ),
                    ),
                    Row(
                      children: [
                        Text(
                          snapshot.data[index]['quantity'],
                          style: const TextStyle(
                            fontWeight: FontWeight.bold,
                            fontSize: 16,
                            fontFamily: 'Montserrat',
                          ),
                        ),
                        Row(
                          children: [
                            Padding(
                              padding: const EdgeInsets.symmetric(horizontal: 5),
                              child: ElevatedButton(onPressed: (){}, child: const Text('+')),
                            ),
                            Padding(
                              padding: const EdgeInsets.all(8.0),
                              child: ElevatedButton(onPressed: (){}, child: const Text('-')),
                            ),
                          ],
                        )
                      ],
                    ),
                  ],
                ),
              );
            },
          );
        } else if (snapshot.hasError) {
          return Center(child: Text('${snapshot.error}'));
        }
        return const Center(child: CircularProgressIndicator());
      },
    );
  }
}

Future<dynamic> getList() async {
  var url = Uri.parse('https://despensa.onrender.com/api/stock');
  var response = await http.get(url);
  if (response.statusCode == 200) {
    return jsonDecode(response.body);
  } else {
    throw Exception('Não foi possível carregar a lista');
  }
}

// Future<dynamic> postItem(
//     String name, String creationDate, String expirationDate) async {
//   var url = Uri.parse('https://despensa.onrender.com/api/stock');
//   var response = await http.post(
//     url,
//     body: json
