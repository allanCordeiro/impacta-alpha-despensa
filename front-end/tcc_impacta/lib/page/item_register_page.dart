import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:tcc_impacta/model/item_model.dart';

import '../provider/item.dart';
import '../repository/implementations/http_api_repository.dart';
import '../routes/app_routes.dart';

class ItemRegister extends StatefulWidget {
  const ItemRegister({super.key});

  @override
  _ItemRegisterState createState() => _ItemRegisterState();
}

class _ItemRegisterState extends State<ItemRegister> {
  final List<ItemModel> _items = [];

  final _formKey = GlobalKey<FormState>();
  final Map<String, dynamic> _formData = {};
  @override
  Widget build(BuildContext context) {
    final HttpApiRepository httpPost = HttpApiRepository();
    final ItemProvider itemProvider = Provider.of(context);
    return Scaffold(
      appBar: AppBar(
        title: const Text('Cadastro'),
      ),
      body: Column(
        children: [
          SizedBox(
            height: 200,
            width: 200,
            child: Image.network(
                'https://cdn.pixabay.com/photo/2015/07/18/09/55/list-850178_1280.jpg'),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: Center(
              child: SizedBox(
                height: 330,
                width: 300,
                child: Container(
                  decoration: const BoxDecoration(
                      color: Color.fromARGB(255, 205, 243, 240),
                      borderRadius: BorderRadius.only(
                          topLeft: Radius.circular(35.0),
                          topRight: Radius.circular(35.0),
                          bottomLeft: Radius.circular(40.0),
                          bottomRight: Radius.circular(40.0))),
                  child: Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: Form(
                      key: _formKey,
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          TextFormField(
                            onSaved: (value) => _formData['name'] = value!,
                            decoration: const InputDecoration(
                              labelText: 'Nome',
                            ),
                          ),
                          TextFormField(
                            onSaved: (value) =>
                                _formData['creationDate'] = value!,
                            decoration: const InputDecoration(
                              labelText: 'Data da compra',
                            ),
                          ),
                          TextFormField(
                            onSaved: (value) => _formData['quantity'] = value!,
                            decoration: const InputDecoration(
                              labelText: 'Quantidade',
                            ),
                            keyboardType: TextInputType.number,
                          ),
                          TextFormField(
                            onSaved: (value) =>
                                _formData['expirationDate'] = value!,
                            // initialValue: _expirationDate.toString(),
                            decoration: const InputDecoration(
                              labelText: 'Data de validade',
                            ),
                          ),
                          const SizedBox(height: 16),
                          Center(
                            child: ElevatedButton(
                              onPressed: () async {
                                _formKey.currentState?.save();
                                await httpPost.addItem(ItemModel(
                                  name: _formData['name'] ?? '',
                                  quantity: _formData['quantity'] ?? '',
                                  creationDate: _formData['creationDate'] ?? '',
                                  expirationDate:
                                      _formData['expirationDate'] ?? '',
                                ));
                                Navigator.of(context)
                                    .pushReplacementNamed(AppRoutes.HOME);
                              },
                              child: const Text('Cadastrar'),
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
